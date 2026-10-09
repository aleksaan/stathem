package models

import (
	"fmt"
	"slices"
	"strings"

	apperrors "github.com/aleksaan/stathem/internal/errors"
)

func (m *Model) Validate() error {
	if err := m.checkEmptyModelName(); err != nil {
		return err
	}
	if err := m.checkEmptyNodesNames(); err != nil {
		return err
	}
	if err := m.checkNodesDuplicates(); err != nil {
		return err
	}
	if err := m.checkMismatchNodes(); err != nil {
		return err
	}
	if err := m.checkNoNodes(); err != nil {
		return err
	}
	// if err := m.checkRulesAreValid(); err != nil {
	// 	return err
	// }
	return nil
}

// func (m *Model) checkRulesAreValid() error {
// 	_, err := rules.New(m.ModelRules)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// ------------------------------------
func (m *Model) checkEmptyModelName() error {

	if len(strings.TrimSpace(m.ModelName)) == 0 {
		return apperrors.ErrEmptyModelName
	}

	return nil
}

// ------------------------------------
func (m *Model) checkEmptyNodesNames() error {

	//there are nodes, but with empty names
	count := 0
	for _, n := range m.ModelNodes {
		if len(strings.TrimSpace(n)) == 0 {
			count++
		}
	}
	if count > 0 {
		return fmt.Errorf("%w: there were found %d nodes with empty names", apperrors.ErrEmptyNodesNames, count)
	}

	return nil
}

// ------------------------------------
func (m *Model) checkNodesDuplicates() error {
	var unique = make(map[string]string)
	var count = 0
	for _, n := range m.ModelNodes {
		if _, ok := unique[n]; ok {
			count++
		}
		unique[n] = n
	}
	if count > 0 {
		return fmt.Errorf("%w:There were found %d nodes duplicates", apperrors.ErrNodesDuplicates, count)
	}
	return nil
}

// ------------------------------------
func (m *Model) checkMismatchNodes() error {
	var count = 0
	for _, e := range m.ModelEdges {
		for n1, n2 := range e {
			if !slices.Contains(m.ModelNodes, n1) {
				count++
			}
			if !slices.Contains(m.ModelNodes, n2) {
				count++
			}
		}
	}
	if count > 0 {
		return fmt.Errorf("%w: There were found %d nodes in edges missing in nodes", apperrors.ErrMismatchingNodes, count)
	}
	return nil
}

// ------------------------------------
func (m *Model) checkNoNodes() error {
	if len(m.ModelNodes) == 0 {
		return apperrors.ErrNoNodes
	}
	return nil
}

func (m *Model) GetPreviousNodes(nodeName string) []string {
	allPreviousNodes := []string{}

	for _, edge := range m.ModelEdges {
		for prevStep, nextStep := range edge {
			if nodeName == nextStep {
				allPreviousNodes = append(allPreviousNodes, prevStep)
			}
		}
	}

	return allPreviousNodes
}

func (m *Model) CheckStepNameIsValid(stepName string) bool {
	res := slices.Contains(m.ModelNodes, stepName)

	return res
}
