pipeline {
    agent any

    parameters {
        choice(
            name: 'OS',
            choices: ['linux', 'darwin', 'windows'],
            description: 'Target operating system'
        )
        choice(
            name: 'ARCH',
            choices: ['arm64', 'amd64'],
            description: 'Target architecture'
        )
        booleanParam(
            name: 'SKIP_TESTS',
            defaultValue: false,
            description: 'Skip running tests'
        )
    }

    environment {
        REPO = "https://github.com/deniszm/obot"
        BRANCH = "develop"
        REGISTRY = "deniszms"
        CGO_ENABLED = "0"
    }

    stages {
        stage('Parameters Info') {
            steps {
                echo "Build Parameters:"
                echo "Target OS: ${params.OS}"
                echo "Target Architecture: ${params.ARCH}"
                echo "Skip Tests: ${params.SKIP_TESTS}"
            }
        }

        stage('Clone') {
            steps {
                echo "Cloning repository ${REPO} on branch ${BRANCH}"
                git branch: "${BRANCH}", url: "${REPO}"
            }
        }

        stage('Test') {
            when {
                expression { 
                    return !params.SKIP_TESTS
                }
            }
            steps {
                echo "Test execution started"
                sh 'make test'
            }
        }

        stage('Build') {
            steps {
                echo "Build execution started for ${params.OS}/${params.ARCH}"
                sh "make TARGETOS=${params.OS} TARGETARCH=${params.ARCH} build"
            }
        }

        stage('Docker Image') {
            steps {
                echo "Image creation started for ${params.OS}/${params.ARCH}"
                sh "make REGISTRY=${REGISTRY} TARGETOS=${params.OS} TARGETARCH=${params.ARCH} image"
            }
        }

        stage('Push to Registry') {
            steps {
                echo "Pushing image to registry ${REGISTRY}"
                script {
                    docker.withRegistry('', 'dockerhub-credentials') {
                        sh "make REGISTRY=${REGISTRY} TARGETOS=${params.OS} TARGETARCH=${params.ARCH} push"
                    }
                }
            }
        }
    }
}