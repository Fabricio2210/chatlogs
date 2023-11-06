pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                // Checkout the code from your GitHub repository
                checkout([$class: 'GitSCM', branches: [[name: 'main']], userRemoteConfigs: [[url: 'https://github.com/Fabricio2210/chatlogs.git']]])
            }
        }

        stage('Build') {
            steps {
                sh 'go get -v' // Download dependencies
                sh 'go build -o myapp' // Build your Go application
            }
        }

        stage('Print Workspace') {
            steps {
                echo "Workspace directory: ${WORKSPACE}"
            }
        }

        stage('Deploy') {
            steps {
                // You can deploy the Go application to your server here
                // For simplicity, we'll just print a message
                echo 'Deploying the application...'
                script {
                    sh 'sudo systemctl restart chatlogs' // Use sudo if needed
                }
            }
        }
    }
}


