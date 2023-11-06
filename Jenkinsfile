pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                // Checkout your source code from the main branch.
                checkout scm
            }
        }

        stage('Install Dependencies') {
            steps {
                sh 'go mod tidy'
            }
        }

        stage('Build') {
            steps {
                // Build your Go application
                sh 'go build -o chatlogs'
            }
        }

        stage('Deploy') {
            steps {
              echo 'done'
            }
        }
    }

    post {
        success {
            echo 'Build and deployment succeeded!'
        }
        failure {
            echo 'Build or deployment failed!'
        }
    }
}

