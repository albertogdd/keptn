package v02

import (
	"encoding/json"
	keptnv2 "github.com/keptn/go-utils/pkg/lib/v0_2_0"
	"github.com/keptn/keptn/shipyard-controller/models"
)

type ModelTransformer struct{}

// TransformToDBModel transforms an instance of models.SequenceExecution to the db specific schema defined in this package
func (ModelTransformer) TransformToDBModel(execution models.SequenceExecution) interface{} {
	return fromSequenceExecution(execution)
}

func (ModelTransformer) TransformToSequenceExecution(dbItem interface{}) (*models.SequenceExecution, error) {
	data, _ := json.Marshal(dbItem)

	internalSequenceExecution := &JsonStringEncodedSequenceExecution{}
	if err := json.Unmarshal(data, internalSequenceExecution); err != nil {
		return nil, err
	}

	// if the current schema version is being used, we need to transform it to model.JsonStringEncodedSequenceExecution
	if internalSequenceExecution.SchemaVersion == SchemaVersionV02 {
		transformedSequenceExecution := internalSequenceExecution.ToSequenceExecution()
		return &transformedSequenceExecution, nil
	}

	// if the old schema is still being used by that item, we can directly unmarshal it to a model.JsonStringEncodedSequenceExecution
	sequenceExecution := &models.SequenceExecution{}
	if err := json.Unmarshal(data, internalSequenceExecution); err != nil {
		return nil, err
	}

	return sequenceExecution, nil
}

func fromSequenceExecution(se models.SequenceExecution) JsonStringEncodedSequenceExecution {
	newSE := JsonStringEncodedSequenceExecution{
		ID: se.ID,
		Sequence: Sequence{
			Name:  se.Sequence.Name,
			Tasks: transformTasks(se.Sequence.Tasks),
		},
		Status:        transformStatus(se.Status),
		Scope:         se.Scope,
		SchemaVersion: SchemaVersionV02,
	}
	if se.InputProperties != nil {
		inputPropertiesJsonString, err := json.Marshal(se.InputProperties)
		if err == nil {
			newSE.InputProperties = string(inputPropertiesJsonString)
		}
	}
	return newSE
}

func transformTasks(tasks []keptnv2.Task) []Task {
	result := []Task{}

	for _, task := range tasks {
		newTask := Task{
			Name:           task.Name,
			TriggeredAfter: task.TriggeredAfter,
		}
		if task.Properties != nil {
			taskPropertiesString, err := json.Marshal(task.Properties)
			if err == nil {
				newTask.Properties = string(taskPropertiesString)
			}
		}
		result = append(result, newTask)
	}
	return result
}

func transformStatus(status models.SequenceExecutionStatus) SequenceExecutionStatus {
	newStatus := SequenceExecutionStatus{
		State:            status.State,
		StateBeforePause: status.StateBeforePause,
		PreviousTasks:    transformPreviousTasks(status.PreviousTasks),
		CurrentTask:      transformCurrentTask(status.CurrentTask),
	}

	return newStatus
}

func transformCurrentTask(task models.TaskExecutionState) TaskExecutionState {
	newTaskExecutionState := TaskExecutionState{
		Name:        task.Name,
		TriggeredID: task.TriggeredID,
		Events:      transformTaskEvents(task.Events),
	}
	return newTaskExecutionState
}

func transformTaskEvents(events []models.TaskEvent) []TaskEvent {
	newTaskEvents := []TaskEvent{}

	for _, e := range events {
		newTaskEvent := TaskEvent{
			EventType: e.EventType,
			Source:    e.Source,
			Result:    e.Result,
			Status:    e.Status,
			Time:      e.Time,
		}

		if e.Properties != nil {
			properties, err := json.Marshal(e.Properties)
			if err == nil {
				newTaskEvent.Properties = string(properties)
			}
		}
		newTaskEvents = append(newTaskEvents, newTaskEvent)
	}
	return newTaskEvents
}

func transformPreviousTasks(tasks []models.TaskExecutionResult) []TaskExecutionResult {
	newPreviousTasks := []TaskExecutionResult{}

	for _, t := range tasks {
		newPreviousTask := TaskExecutionResult{
			Name:        t.Name,
			TriggeredID: t.TriggeredID,
			Result:      t.Result,
			Status:      t.Status,
		}

		if t.Properties != nil {
			properties, err := json.Marshal(t.Properties)
			if err == nil {
				newPreviousTask.Properties = string(properties)
			}
		}
		newPreviousTasks = append(newPreviousTasks, newPreviousTask)
	}
	return newPreviousTasks
}
