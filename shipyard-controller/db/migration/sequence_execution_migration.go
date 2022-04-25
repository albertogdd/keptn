package migration

import (
	"fmt"
	apimodels "github.com/keptn/go-utils/pkg/api/models"
	keptnv2 "github.com/keptn/go-utils/pkg/lib/v0_2_0"
	"github.com/keptn/keptn/shipyard-controller/db"
	"github.com/keptn/keptn/shipyard-controller/models"
	logger "github.com/sirupsen/logrus"
)

// NewSequenceExecutionMigrator creates a new SequenceExecutionMigrator
// Internally it is using the SequenceExecutionJsonStringRepo decorator
// which stores the arbitrary event payload sent by keptn integrations as Json strings to avoid having property names with dots (.) in them
func NewSequenceExecutionMigrator(dbConnection *db.MongoDBConnection) *SequenceExecutionMigrator {
	return &SequenceExecutionMigrator{
		projectRepo:           db.NewMongoDBKeyEncodingProjectsRepo(dbConnection),
		sequenceExecutionRepo: db.NewMongoDBSequenceExecutionRepo(dbConnection),
	}
}

type SequenceExecutionMigrator struct {
	sequenceExecutionRepo db.SequenceExecutionRepo
	projectRepo           db.ProjectRepo
}

// MigrateSequenceExecutions retrieves all existing sequence executions from the repository
// and performs an update operation on each of them using the SequenceExecutionJsonStringRepo.
// This way, sequence executions containing stored with the previous format are migrated to the new one
func (s *SequenceExecutionMigrator) MigrateSequenceExecutions() error {
	projects, err := s.projectRepo.GetProjects()
	if err != nil {
		return fmt.Errorf("could not migrate sequence executions to new format: %w", err)
	}
	return s.updateSequenceExecutionsOfProject(projects)
}

func (s *SequenceExecutionMigrator) updateSequenceExecutionsOfProject(projects []*apimodels.ExpandedProject) error {
	if projects == nil {
		return nil
	}
	for _, project := range projects {
		sequenceExecutions, err := s.sequenceExecutionRepo.Get(models.SequenceExecutionFilter{Scope: models.EventScope{
			EventData: keptnv2.EventData{
				Project: project.ProjectName,
			},
		}})
		if err != nil {
			logger.Errorf("Could not retrieve sequence executions for project %s: %v", project.ProjectName, err)
			continue
		}
		for _, sequenceExecution := range sequenceExecutions {
			// check if sequence execution has already been migrated
			if err := s.sequenceExecutionRepo.Upsert(sequenceExecution, nil); err != nil {
				logger.Errorf("Could not update sequence execution with ID %s for project %s: %v", sequenceExecution.ID, project.ProjectName, err)
				continue
			}
		}
	}
	return nil
}