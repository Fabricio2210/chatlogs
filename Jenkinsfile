pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
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
                sh 'go build -o chatlogs'
            }
        }

        stage('Deploy') {
            steps {
                // Start the new Go executable
                sh 'nohup ./chatlogs > /dev/null 2>&1 &'
                echo $! > /var/run/chatlogs.pid
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

