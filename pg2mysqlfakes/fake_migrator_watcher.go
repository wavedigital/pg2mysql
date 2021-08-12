// This file was generated by counterfeiter
package pg2mysqlfakes

import (
	"sync"

	"github.com/wavedigital/pg2mysql"
)

type FakeMigratorWatcher struct {
	WillBuildSchemaStub                          func()
	willBuildSchemaMutex                         sync.RWMutex
	willBuildSchemaArgsForCall                   []struct{}
	DidBuildSchemaStub                           func()
	didBuildSchemaMutex                          sync.RWMutex
	didBuildSchemaArgsForCall                    []struct{}
	WillDisableConstraintsStub                   func()
	willDisableConstraintsMutex                  sync.RWMutex
	willDisableConstraintsArgsForCall            []struct{}
	DidDisableConstraintsStub                    func()
	didDisableConstraintsMutex                   sync.RWMutex
	didDisableConstraintsArgsForCall             []struct{}
	WillEnableConstraintsStub                    func()
	willEnableConstraintsMutex                   sync.RWMutex
	willEnableConstraintsArgsForCall             []struct{}
	EnableConstraintsDidFinishStub               func()
	enableConstraintsDidFinishMutex              sync.RWMutex
	enableConstraintsDidFinishArgsForCall        []struct{}
	EnableConstraintsDidFailWithErrorStub        func(err error)
	enableConstraintsDidFailWithErrorMutex       sync.RWMutex
	enableConstraintsDidFailWithErrorArgsForCall []struct {
		err error
	}
	WillTruncateTableStub        func(tableName string)
	willTruncateTableMutex       sync.RWMutex
	willTruncateTableArgsForCall []struct {
		tableName string
	}
	TruncateTableDidFinishStub        func(tableName string)
	truncateTableDidFinishMutex       sync.RWMutex
	truncateTableDidFinishArgsForCall []struct {
		tableName string
	}
	TableMigrationDidStartStub        func(tableName string)
	tableMigrationDidStartMutex       sync.RWMutex
	tableMigrationDidStartArgsForCall []struct {
		tableName string
	}
	TableMigrationDidFinishStub        func(tableName string, recordsInserted int64)
	tableMigrationDidFinishMutex       sync.RWMutex
	tableMigrationDidFinishArgsForCall []struct {
		tableName       string
		recordsInserted int64
	}
	DidMigrateRowStub        func(tableName string)
	didMigrateRowMutex       sync.RWMutex
	didMigrateRowArgsForCall []struct {
		tableName string
	}
	DidFailToMigrateRowWithErrorStub        func(tableName string, err error)
	didFailToMigrateRowWithErrorMutex       sync.RWMutex
	didFailToMigrateRowWithErrorArgsForCall []struct {
		tableName string
		err       error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMigratorWatcher) WillBuildSchema() {
	fake.willBuildSchemaMutex.Lock()
	fake.willBuildSchemaArgsForCall = append(fake.willBuildSchemaArgsForCall, struct{}{})
	fake.recordInvocation("WillBuildSchema", []interface{}{})
	fake.willBuildSchemaMutex.Unlock()
	if fake.WillBuildSchemaStub != nil {
		fake.WillBuildSchemaStub()
	}
}

func (fake *FakeMigratorWatcher) WillBuildSchemaCallCount() int {
	fake.willBuildSchemaMutex.RLock()
	defer fake.willBuildSchemaMutex.RUnlock()
	return len(fake.willBuildSchemaArgsForCall)
}

func (fake *FakeMigratorWatcher) DidBuildSchema() {
	fake.didBuildSchemaMutex.Lock()
	fake.didBuildSchemaArgsForCall = append(fake.didBuildSchemaArgsForCall, struct{}{})
	fake.recordInvocation("DidBuildSchema", []interface{}{})
	fake.didBuildSchemaMutex.Unlock()
	if fake.DidBuildSchemaStub != nil {
		fake.DidBuildSchemaStub()
	}
}

func (fake *FakeMigratorWatcher) DidBuildSchemaCallCount() int {
	fake.didBuildSchemaMutex.RLock()
	defer fake.didBuildSchemaMutex.RUnlock()
	return len(fake.didBuildSchemaArgsForCall)
}

func (fake *FakeMigratorWatcher) WillDisableConstraints() {
	fake.willDisableConstraintsMutex.Lock()
	fake.willDisableConstraintsArgsForCall = append(fake.willDisableConstraintsArgsForCall, struct{}{})
	fake.recordInvocation("WillDisableConstraints", []interface{}{})
	fake.willDisableConstraintsMutex.Unlock()
	if fake.WillDisableConstraintsStub != nil {
		fake.WillDisableConstraintsStub()
	}
}

func (fake *FakeMigratorWatcher) WillDisableConstraintsCallCount() int {
	fake.willDisableConstraintsMutex.RLock()
	defer fake.willDisableConstraintsMutex.RUnlock()
	return len(fake.willDisableConstraintsArgsForCall)
}

func (fake *FakeMigratorWatcher) DidDisableConstraints() {
	fake.didDisableConstraintsMutex.Lock()
	fake.didDisableConstraintsArgsForCall = append(fake.didDisableConstraintsArgsForCall, struct{}{})
	fake.recordInvocation("DidDisableConstraints", []interface{}{})
	fake.didDisableConstraintsMutex.Unlock()
	if fake.DidDisableConstraintsStub != nil {
		fake.DidDisableConstraintsStub()
	}
}

func (fake *FakeMigratorWatcher) DidDisableConstraintsCallCount() int {
	fake.didDisableConstraintsMutex.RLock()
	defer fake.didDisableConstraintsMutex.RUnlock()
	return len(fake.didDisableConstraintsArgsForCall)
}

func (fake *FakeMigratorWatcher) WillEnableConstraints() {
	fake.willEnableConstraintsMutex.Lock()
	fake.willEnableConstraintsArgsForCall = append(fake.willEnableConstraintsArgsForCall, struct{}{})
	fake.recordInvocation("WillEnableConstraints", []interface{}{})
	fake.willEnableConstraintsMutex.Unlock()
	if fake.WillEnableConstraintsStub != nil {
		fake.WillEnableConstraintsStub()
	}
}

func (fake *FakeMigratorWatcher) WillEnableConstraintsCallCount() int {
	fake.willEnableConstraintsMutex.RLock()
	defer fake.willEnableConstraintsMutex.RUnlock()
	return len(fake.willEnableConstraintsArgsForCall)
}

func (fake *FakeMigratorWatcher) EnableConstraintsDidFinish() {
	fake.enableConstraintsDidFinishMutex.Lock()
	fake.enableConstraintsDidFinishArgsForCall = append(fake.enableConstraintsDidFinishArgsForCall, struct{}{})
	fake.recordInvocation("EnableConstraintsDidFinish", []interface{}{})
	fake.enableConstraintsDidFinishMutex.Unlock()
	if fake.EnableConstraintsDidFinishStub != nil {
		fake.EnableConstraintsDidFinishStub()
	}
}

func (fake *FakeMigratorWatcher) EnableConstraintsDidFinishCallCount() int {
	fake.enableConstraintsDidFinishMutex.RLock()
	defer fake.enableConstraintsDidFinishMutex.RUnlock()
	return len(fake.enableConstraintsDidFinishArgsForCall)
}

func (fake *FakeMigratorWatcher) EnableConstraintsDidFailWithError(err error) {
	fake.enableConstraintsDidFailWithErrorMutex.Lock()
	fake.enableConstraintsDidFailWithErrorArgsForCall = append(fake.enableConstraintsDidFailWithErrorArgsForCall, struct {
		err error
	}{err})
	fake.recordInvocation("EnableConstraintsDidFailWithError", []interface{}{err})
	fake.enableConstraintsDidFailWithErrorMutex.Unlock()
	if fake.EnableConstraintsDidFailWithErrorStub != nil {
		fake.EnableConstraintsDidFailWithErrorStub(err)
	}
}

func (fake *FakeMigratorWatcher) EnableConstraintsDidFailWithErrorCallCount() int {
	fake.enableConstraintsDidFailWithErrorMutex.RLock()
	defer fake.enableConstraintsDidFailWithErrorMutex.RUnlock()
	return len(fake.enableConstraintsDidFailWithErrorArgsForCall)
}

func (fake *FakeMigratorWatcher) EnableConstraintsDidFailWithErrorArgsForCall(i int) error {
	fake.enableConstraintsDidFailWithErrorMutex.RLock()
	defer fake.enableConstraintsDidFailWithErrorMutex.RUnlock()
	return fake.enableConstraintsDidFailWithErrorArgsForCall[i].err
}

func (fake *FakeMigratorWatcher) WillTruncateTable(tableName string) {
	fake.willTruncateTableMutex.Lock()
	fake.willTruncateTableArgsForCall = append(fake.willTruncateTableArgsForCall, struct {
		tableName string
	}{tableName})
	fake.recordInvocation("WillTruncateTable", []interface{}{tableName})
	fake.willTruncateTableMutex.Unlock()
	if fake.WillTruncateTableStub != nil {
		fake.WillTruncateTableStub(tableName)
	}
}

func (fake *FakeMigratorWatcher) WillTruncateTableCallCount() int {
	fake.willTruncateTableMutex.RLock()
	defer fake.willTruncateTableMutex.RUnlock()
	return len(fake.willTruncateTableArgsForCall)
}

func (fake *FakeMigratorWatcher) WillTruncateTableArgsForCall(i int) string {
	fake.willTruncateTableMutex.RLock()
	defer fake.willTruncateTableMutex.RUnlock()
	return fake.willTruncateTableArgsForCall[i].tableName
}

func (fake *FakeMigratorWatcher) TruncateTableDidFinish(tableName string) {
	fake.truncateTableDidFinishMutex.Lock()
	fake.truncateTableDidFinishArgsForCall = append(fake.truncateTableDidFinishArgsForCall, struct {
		tableName string
	}{tableName})
	fake.recordInvocation("TruncateTableDidFinish", []interface{}{tableName})
	fake.truncateTableDidFinishMutex.Unlock()
	if fake.TruncateTableDidFinishStub != nil {
		fake.TruncateTableDidFinishStub(tableName)
	}
}

func (fake *FakeMigratorWatcher) TruncateTableDidFinishCallCount() int {
	fake.truncateTableDidFinishMutex.RLock()
	defer fake.truncateTableDidFinishMutex.RUnlock()
	return len(fake.truncateTableDidFinishArgsForCall)
}

func (fake *FakeMigratorWatcher) TruncateTableDidFinishArgsForCall(i int) string {
	fake.truncateTableDidFinishMutex.RLock()
	defer fake.truncateTableDidFinishMutex.RUnlock()
	return fake.truncateTableDidFinishArgsForCall[i].tableName
}

func (fake *FakeMigratorWatcher) TableMigrationDidStart(tableName string) {
	fake.tableMigrationDidStartMutex.Lock()
	fake.tableMigrationDidStartArgsForCall = append(fake.tableMigrationDidStartArgsForCall, struct {
		tableName string
	}{tableName})
	fake.recordInvocation("TableMigrationDidStart", []interface{}{tableName})
	fake.tableMigrationDidStartMutex.Unlock()
	if fake.TableMigrationDidStartStub != nil {
		fake.TableMigrationDidStartStub(tableName)
	}
}

func (fake *FakeMigratorWatcher) TableMigrationDidStartCallCount() int {
	fake.tableMigrationDidStartMutex.RLock()
	defer fake.tableMigrationDidStartMutex.RUnlock()
	return len(fake.tableMigrationDidStartArgsForCall)
}

func (fake *FakeMigratorWatcher) TableMigrationDidStartArgsForCall(i int) string {
	fake.tableMigrationDidStartMutex.RLock()
	defer fake.tableMigrationDidStartMutex.RUnlock()
	return fake.tableMigrationDidStartArgsForCall[i].tableName
}

func (fake *FakeMigratorWatcher) TableMigrationDidFinish(tableName string, recordsInserted int64) {
	fake.tableMigrationDidFinishMutex.Lock()
	fake.tableMigrationDidFinishArgsForCall = append(fake.tableMigrationDidFinishArgsForCall, struct {
		tableName       string
		recordsInserted int64
	}{tableName, recordsInserted})
	fake.recordInvocation("TableMigrationDidFinish", []interface{}{tableName, recordsInserted})
	fake.tableMigrationDidFinishMutex.Unlock()
	if fake.TableMigrationDidFinishStub != nil {
		fake.TableMigrationDidFinishStub(tableName, recordsInserted)
	}
}

func (fake *FakeMigratorWatcher) TableMigrationDidFinishCallCount() int {
	fake.tableMigrationDidFinishMutex.RLock()
	defer fake.tableMigrationDidFinishMutex.RUnlock()
	return len(fake.tableMigrationDidFinishArgsForCall)
}

func (fake *FakeMigratorWatcher) TableMigrationDidFinishArgsForCall(i int) (string, int64) {
	fake.tableMigrationDidFinishMutex.RLock()
	defer fake.tableMigrationDidFinishMutex.RUnlock()
	return fake.tableMigrationDidFinishArgsForCall[i].tableName, fake.tableMigrationDidFinishArgsForCall[i].recordsInserted
}

func (fake *FakeMigratorWatcher) DidMigrateRow(tableName string) {
	fake.didMigrateRowMutex.Lock()
	fake.didMigrateRowArgsForCall = append(fake.didMigrateRowArgsForCall, struct {
		tableName string
	}{tableName})
	fake.recordInvocation("DidMigrateRow", []interface{}{tableName})
	fake.didMigrateRowMutex.Unlock()
	if fake.DidMigrateRowStub != nil {
		fake.DidMigrateRowStub(tableName)
	}
}

func (fake *FakeMigratorWatcher) DidMigrateRowCallCount() int {
	fake.didMigrateRowMutex.RLock()
	defer fake.didMigrateRowMutex.RUnlock()
	return len(fake.didMigrateRowArgsForCall)
}

func (fake *FakeMigratorWatcher) DidMigrateRowArgsForCall(i int) string {
	fake.didMigrateRowMutex.RLock()
	defer fake.didMigrateRowMutex.RUnlock()
	return fake.didMigrateRowArgsForCall[i].tableName
}

func (fake *FakeMigratorWatcher) DidFailToMigrateRowWithError(tableName string, err error) {
	fake.didFailToMigrateRowWithErrorMutex.Lock()
	fake.didFailToMigrateRowWithErrorArgsForCall = append(fake.didFailToMigrateRowWithErrorArgsForCall, struct {
		tableName string
		err       error
	}{tableName, err})
	fake.recordInvocation("DidFailToMigrateRowWithError", []interface{}{tableName, err})
	fake.didFailToMigrateRowWithErrorMutex.Unlock()
	if fake.DidFailToMigrateRowWithErrorStub != nil {
		fake.DidFailToMigrateRowWithErrorStub(tableName, err)
	}
}

func (fake *FakeMigratorWatcher) DidFailToMigrateRowWithErrorCallCount() int {
	fake.didFailToMigrateRowWithErrorMutex.RLock()
	defer fake.didFailToMigrateRowWithErrorMutex.RUnlock()
	return len(fake.didFailToMigrateRowWithErrorArgsForCall)
}

func (fake *FakeMigratorWatcher) DidFailToMigrateRowWithErrorArgsForCall(i int) (string, error) {
	fake.didFailToMigrateRowWithErrorMutex.RLock()
	defer fake.didFailToMigrateRowWithErrorMutex.RUnlock()
	return fake.didFailToMigrateRowWithErrorArgsForCall[i].tableName, fake.didFailToMigrateRowWithErrorArgsForCall[i].err
}

func (fake *FakeMigratorWatcher) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.willBuildSchemaMutex.RLock()
	defer fake.willBuildSchemaMutex.RUnlock()
	fake.didBuildSchemaMutex.RLock()
	defer fake.didBuildSchemaMutex.RUnlock()
	fake.willDisableConstraintsMutex.RLock()
	defer fake.willDisableConstraintsMutex.RUnlock()
	fake.didDisableConstraintsMutex.RLock()
	defer fake.didDisableConstraintsMutex.RUnlock()
	fake.willEnableConstraintsMutex.RLock()
	defer fake.willEnableConstraintsMutex.RUnlock()
	fake.enableConstraintsDidFinishMutex.RLock()
	defer fake.enableConstraintsDidFinishMutex.RUnlock()
	fake.enableConstraintsDidFailWithErrorMutex.RLock()
	defer fake.enableConstraintsDidFailWithErrorMutex.RUnlock()
	fake.willTruncateTableMutex.RLock()
	defer fake.willTruncateTableMutex.RUnlock()
	fake.truncateTableDidFinishMutex.RLock()
	defer fake.truncateTableDidFinishMutex.RUnlock()
	fake.tableMigrationDidStartMutex.RLock()
	defer fake.tableMigrationDidStartMutex.RUnlock()
	fake.tableMigrationDidFinishMutex.RLock()
	defer fake.tableMigrationDidFinishMutex.RUnlock()
	fake.didMigrateRowMutex.RLock()
	defer fake.didMigrateRowMutex.RUnlock()
	fake.didFailToMigrateRowWithErrorMutex.RLock()
	defer fake.didFailToMigrateRowWithErrorMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeMigratorWatcher) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ pg2mysql.MigratorWatcher = new(FakeMigratorWatcher)
