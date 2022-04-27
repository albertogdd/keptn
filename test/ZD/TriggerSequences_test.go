package ZD

import (
	"fmt"
	"github.com/keptn/go-utils/pkg/api/models"
	"github.com/keptn/go-utils/pkg/lib/keptn"
	"github.com/keptn/go-utils/pkg/lib/v0_2_0"
	testutils "github.com/keptn/keptn/test/go-tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v3"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type TestSuiteSequences struct {
	suite.Suite
	env     *ZeroDowntimeEnv
	project string
}

type TriggeredSequence struct {
	keptnContext string
	projectName  string
	sequenceName string
}

func NewTriggeredSequence(keptnContext string, projectName string, seqName string) TriggeredSequence {
	return TriggeredSequence{
		keptnContext: keptnContext,
		projectName:  projectName,
		sequenceName: seqName,
	}
}

func (suite *TestSuiteSequences) SetupSuite() {

	suite.T().Log("Starting test for sequences")
	// setup a project every clock ticker
	suite.createNew()
}

func (suite *TestSuiteSequences) createNew() {
	var err error
	projectName := "zd" + suite.gedId()
	suite.T().Logf("Creating project with id %s ", projectName)
	suite.project, err = testutils.CreateProject(projectName, suite.env.ShipyardFile)
	suite.Nil(err)
	output, err := testutils.ExecuteCommand(fmt.Sprintf("keptn create service %s --project=%s", "myservice", suite.project))
	suite.Nil(err)
	suite.Contains(output, "created successfully")

	suite.T().Logf("Starting test for project %s ", suite.project)
}

func (suite *TestSuiteSequences) BeforeTest(suiteName, testName string) {
	atomic.AddUint64(&suite.env.FiredSequences, 1)
	suite.T().Log("Running one more test, tot ", suite.env.FiredSequences)
}

//Test_Sequences can be used to test a single run of the test suite
func Test_Sequences(t *testing.T) {
	////TODO setup a one run env

	Env := SetupZD()
	var err error
	Env.ExistingProject, err = testutils.CreateProject("projectzd", Env.ShipyardFile)
	assert.Nil(t, err)
	_, err = testutils.ExecuteCommand(fmt.Sprintf("keptn create service %s --project=%s", "myservice", Env.ExistingProject))
	assert.Nil(t, err)

	s := &TestSuiteSequences{
		env: SetupZD(),
	}
	suite.Run(t, s)
}

// to perform tests sequentially inside ZD
func Sequences(t *testing.T, env *ZeroDowntimeEnv) {
	var s *TestSuiteSequences
	env.Wg.Add(1)
	wgSequences := &sync.WaitGroup{}
Loop:
	for {
		select {
		case <-env.Ctx.Done():
			break Loop
		case <-env.SeqTicker.C:
			s = &TestSuiteSequences{
				env: env,
			}
			wgSequences.Add(1)
			go func() {
				suite.Run(t, s)
				wgSequences.Done()
			}()

		}
	}

	wgSequences.Wait()
	t.Run("Summary", func(t2 *testing.T) {
		PrintSequencesResults(t2, env)
	})
	env.Wg.Done()

}

//webhook

//trigger a sequence while graceful

//approval

//func Test_Sequences(t *testing.T) {
//	suite.Run(t, new(TestSuiteSequences))
//}

func (suite *TestSuiteSequences) Test_EvaluationFails() {
	suite.T().Log("Started Ev")
	suite.trigger("evaluation", nil, false)
}

func (suite *TestSuiteSequences) Test_DeliveryFails() {
	suite.trigger("delivery", nil, false)
}

func (suite *TestSuiteSequences) Test_ExistingEvaluationFails() {
	suite.trigger("evaluation", nil, true)
}

func (suite *TestSuiteSequences) Test_ExistingDeliveryFails() {
	suite.trigger("delivery", nil, true)
}

func (suite *TestSuiteSequences) trigger(triggerType string, data keptn.EventProperties, existing bool) {
	suite.T().Log("Started Trigger")
	project := suite.project
	if existing {
		project = suite.env.ExistingProject
	}

	suite.T().Logf("triggering sequence %s for project %s", triggerType, project)
	// trigger a delivery sequence
	keptnContext, err := testutils.TriggerSequence(project, "myservice", "dev", triggerType, data)
	suite.Nil(err)
	suite.T().Logf("triggered sequence %s for project %s with context %s", triggerType, project, keptnContext)
	sequence := NewTriggeredSequence(keptnContext, project, triggerType)
	//sequences.Add(sequence)

	suite.checkSequence(sequence)
}

func (suite *TestSuiteSequences) checkSequence(sequence TriggeredSequence) {

	var sequenceFinishedEvent *models.KeptnContextExtendedCE
	stageSequenceName := fmt.Sprintf("%s.%s", "dev", sequence.sequenceName)
	var err error

	suite.T().Logf("verifying completion of sequence %s with keptnContext %s in project %s", sequence.sequenceName, sequence.keptnContext, sequence.projectName)
	suite.Eventually(func() bool {
		sequenceFinishedEvent, err = testutils.GetLatestEventOfType(sequence.keptnContext, sequence.projectName, "dev", v0_2_0.GetFinishedEventType(stageSequenceName))
		if sequenceFinishedEvent == nil || err != nil {
			return false
		}
		atomic.AddUint64(&suite.env.PassedSequences, 1)
		return true
	}, 15*time.Second, 5*time.Second)

	if sequenceFinishedEvent == nil || err != nil {
		atomic.AddUint64(&suite.env.FailedSequences, 1)
		suite.T().Errorf("sequence %s with keptnContext %s in project %s has NOT been finished", sequence.sequenceName, sequence.keptnContext, sequence.projectName)

	} else {
		suite.T().Logf("sequence %s with keptnContext %s in project %s has been finished", sequence.sequenceName, sequence.keptnContext, sequence.projectName)
	}
}

func GetShipyard() (string, error) {
	shipyard := &v0_2_0.Shipyard{
		ApiVersion: "0.2.3",
		Kind:       "shipyard",
		Metadata:   v0_2_0.Metadata{},
		Spec: v0_2_0.ShipyardSpec{
			Stages: []v0_2_0.Stage{},
		},
	}

	stage := v0_2_0.Stage{
		Name: "dev",
		Sequences: []v0_2_0.Sequence{
			{
				Name: "hooks",
				Tasks: []v0_2_0.Task{
					{
						Name: "mytask",
					},
				},
			},
		},
	}

	shipyard.Spec.Stages = append(shipyard.Spec.Stages, stage)

	shipyardFileContent, _ := yaml.Marshal(shipyard)

	return testutils.CreateTmpShipyardFile(string(shipyardFileContent))
}

func (suite *TestSuiteSequences) gedId() string {
	atomic.AddUint64(&suite.env.Id, 1)
	return fmt.Sprintf("%d", suite.env.Id)
}

func PrintSequencesResults(t *testing.T, env *ZeroDowntimeEnv) {

	t.Log("-----------------------------------------------")
	t.Log("TOTAL SEQUENCES: ", env.FiredSequences)
	t.Log("TOTAL SUCCESS ", env.PassedSequences)
	t.Log("TOTAL FAILURES ", env.FailedSequences)
	t.Log("-----------------------------------------------")

}
