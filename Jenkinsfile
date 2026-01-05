pipeline {
    agent any

    environment {
        IMAGE_NAME = "go-db-app"
        IMAGE_TAG  = "${BUILD_NUMBER}"
    }

    stages {

        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Go Build & Test (Docker)') {
            steps {
                script {
                    docker.image('golang:1.24.11-alpine3.23').inside('-e GOCACHE=/tmp/go-cache') {
    sh '''
      cd go-db-app
      go version
      go env GOCACHE
      go mod download
      go test ./...
      go build -o app
    '''
}
                }
            }
        }

        stage('Docker Build') {
            steps {
                script {
                    docker.build(
                        "${IMAGE_NAME}:${IMAGE_TAG}",
                        "-f Docker/Dockerfile go-db-app"
                    )
                }
            }
        }

        stage('Smoke Test') {
            steps {
                sh '''
                  docker run --rm ${IMAGE_NAME}:${IMAGE_TAG} ./app --help || true
                '''
            }
        }

        // Optional push stage
        stage('Push Image') {
            when {
                branch 'main'
            }
            steps {
                script {
                    docker.withRegistry('', 'dockerhub-creds') {
                        docker.image("${IMAGE_NAME}:${IMAGE_TAG}").push()
                        docker.image("${IMAGE_NAME}:${IMAGE_TAG}").push('latest')
                    }
                }
            }
        }
    }
}
