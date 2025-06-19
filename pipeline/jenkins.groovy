pipeline {
    agent any

    environment {
        REPO = "https://github.com/deniszm/obot"
        BRANCH = "main"
        REGISTRY = "deniszms"
    }

    stages {
        stage('clone') {
            steps {
                echo "Cloning repository ${REPO} on branch ${BRANCH}"
                git branch: "${BRANCH}", url: "${REPO}"
            }
        }

        stage('test') {
            steps {
                echo "Test execution started"
                sh 'make test'
            }
        }

        stage('build') {
            steps {
                echo "Build execution started"
                sh 'make build'
            }
        }

        stage('image') {
            steps {
                echo "Image creation started"
                sh 'make REGISTRY=${REGISTRY} image'
            }
        }

        stage('push') {
            steps {
                echo "Pushing image to registry"
                script {
                    docker.withRegistry('', 'dockerhub-credentials') {
                        sh "make REGISTRY=${REGISTRY} push"
                    }
                }
            }
        }
    }
}