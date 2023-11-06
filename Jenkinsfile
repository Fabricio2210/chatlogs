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

        stage('Prepare Script') {
            steps {
                // Add execute permission to the script
                sh 'chmod +x run.sh'
            }
        }

        stage('Deploy') {
            steps {
                script {
                    sh 'nohup ./run.sh > /dev/null 2>&1 &'
                }
            }
        }
    }
    stage('Debug') {
        steps {
            sh 'pwd'
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


