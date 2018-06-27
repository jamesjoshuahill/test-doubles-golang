// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/jamesjoshuahill/test-doubles-golang (interfaces: repository)

package pegomock

import (
	test_doubles_golang "github.com/jamesjoshuahill/test-doubles-golang"
	pegomock "github.com/petergtz/pegomock"
	"reflect"
)

type Mockrepository struct {
	fail func(message string, callerSkip ...int)
}

func NewMockrepository() *Mockrepository {
	return &Mockrepository{fail: pegomock.GlobalFailHandler}
}

func (mock *Mockrepository) Query(name string, kind string) ([]test_doubles_golang.Record, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockMockrepository().")
	}
	params := []pegomock.Param{name, kind}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Query", params, []reflect.Type{reflect.TypeOf((*[]test_doubles_golang.Record)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 []test_doubles_golang.Record
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].([]test_doubles_golang.Record)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *Mockrepository) VerifyWasCalledOnce() *Verifierrepository {
	return &Verifierrepository{mock, pegomock.Times(1), nil}
}

func (mock *Mockrepository) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *Verifierrepository {
	return &Verifierrepository{mock, invocationCountMatcher, nil}
}

func (mock *Mockrepository) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *Verifierrepository {
	return &Verifierrepository{mock, invocationCountMatcher, inOrderContext}
}

type Verifierrepository struct {
	mock                   *Mockrepository
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
}

func (verifier *Verifierrepository) Query(name string, kind string) *repository_Query_OngoingVerification {
	params := []pegomock.Param{name, kind}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Query", params)
	return &repository_Query_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type repository_Query_OngoingVerification struct {
	mock              *Mockrepository
	methodInvocations []pegomock.MethodInvocation
}

func (c *repository_Query_OngoingVerification) GetCapturedArguments() (string, string) {
	name, kind := c.GetAllCapturedArguments()
	return name[len(name)-1], kind[len(kind)-1]
}

func (c *repository_Query_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]string, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
	}
	return
}
