pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                // Checkout your source code from a version control system like Git.
                checkout scm
            }
        }
        stage('Install Dependencies') {
            steps {
                // Ensure dependencies are up to date using go mod tidy
                sh 'go mod tidy'
            }
        }
        stage('Build') {
            steps {
                // Build your Go application
                sh 'go build -o chatlogs'
            }
        }


        // stage('Deploy') {
        //     steps {
        //         // Deploy your Go application
        //         sh './chatlogs'
        //     }
        // }
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

