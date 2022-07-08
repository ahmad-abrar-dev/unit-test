package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// === unit test required ===
// the file name must contain "_test" in the last word
// the function name must contain "Test" in the first word
// the argument inside the function test must contain pckg (t *testing.T)

// to do the testing, go to the path and run the command below
// >> go test //it will run all file within _test

//benchmark

func TestMain(m *testing.M) {
	fmt.Println("====Before Unit Test")

	m.Run() //running the testing

	fmt.Println("====After Unit Test")
}

func TestHelloWorldZhun(t *testing.T) {
	result := HelloWorld("Zhun")
	if result != "Hello Zhun" {
		//unit test failed
		// panic("Result is not Hello Zhun") //use panic isn't recommended way in unit test
		// t.FailNow() // same as below code but no message
		t.Fatal("Result Must Be Hello Zhun")
	}
	fmt.Println("Test Zhun Done")
}

func TestHelloWorldWeky(t *testing.T) {
	result := HelloWorld("Weky")
	if result != "Hello Weky" {
		//unit test failed
		// panic("Result is not Hello Weky") //use panic isn't recommended way in unit test
		// t.Fail() // same as below code but no message
		t.Error("Result Must Be Hello Weky")
	}
	fmt.Println("Test Weky Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Zhun")
	// the method Equal same as fail() in assert pkg
	assert.Equal(t, "Hello Zhun", result, "Result doesn't contain Hello Zhun") //params : unitTestGoPkg, ValueExpactating, parseValue, ErrMessage
	fmt.Println("TestHelloWorld with assert done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Zhun")
	// the method Equal same as failNow() in require pkg
	require.Equal(t, "Hello Zhun", result, "Result doesn't contain Hello Zhun") //params : unitTestGoPkg, ValueExpactating, parseValue, ErrMessage
	fmt.Println("TestHelloWorld with require done")                             // just to show if the code below require package doesn't work
}

func TestSkip(t *testing.T) {
	// darwin to read the mac os
	if runtime.GOOS == "windows" {
		t.Skip("Can not run on Windows")
	}
	result := HelloWorld("Zhun")
	// the method Equal same as failNow() in require pkg
	require.Equal(t, "Hello Zhun", result, "Result doesn't contain Hello Zhun")
}

func TestSubTest(t *testing.T) {
	//to run specific sub test use : go test -v -run=<funcTest>/<subtest name>
	t.Run("zhun", func(t *testing.T) {
		result := HelloWorld("zhun")
		require.Equal(t, "Hello zhun", result)
	})
	t.Run("weky", func(t *testing.T) {
		result := HelloWorld("weky")
		require.Equal(t, "Hi weky", result)
	})
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "testDongHelloWorld(zhun)",
			request:  "zhun",
			expected: "Hello zhun",
		},
		{
			name:     "testDongHelloWorld(weky)",
			request:  "weky",
			expected: "Hello weky",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}

}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Zhun")
	}
}

func BenchmarkHelloWorlda(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Zhun")
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Zhun", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Zhun")
		}
	})

	b.Run("Weky", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Weky")
		}
	})
}

func BenchmarkHelloWorldTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "testDongHelloWorld(zhun)",
			request: "zhun",
		},
		{
			name:    "testDongHelloWorld(weky)",
			request: "weky",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}

}
