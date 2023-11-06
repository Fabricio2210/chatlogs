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
                // Restart the Systemd service to pick up the changes
                sh 'ssh fabricio@127.0.0.1 sudo systemctl restart chatlogs.service'
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

