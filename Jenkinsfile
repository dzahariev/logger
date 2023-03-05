pipeline {
    // install golang 1.20.1 on Jenkins node
    agent any
    tools {
        go 'go1.20.1'
    }
    environment {
        CGO_ENABLED = 0 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {
        stage("clean") {
            steps {
                echo 'Clean step started ...'
                sh '''go clean
                rm -f coverage.out
                rm -fr bin'''
            }
        }

        stage("format") {
            steps {
                echo 'Format step started ...'
                sh 'go fmt ./...'
            }
        }
        
        stage("build") {
            steps {
                echo 'Build step started ...'
                sh 'go build -o bin/logger'
            }
        }
        
        stage("tests") {
            steps {
                echo 'Tests step started ...'
                sh 'go test -v -coverpkg $(shell go list ./... | egrep -v "test" | paste -sd "," -) ./... -coverprofile=coverage.out -covermode=atomic'
            }
        }
    }
}
