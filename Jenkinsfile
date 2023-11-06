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

        stage('Deploy') {
            steps {
                script {
                    // Check if a previous script is running and terminate it
                    def pidFile = sh(script: 'cat /var/run/run.pid', returnStatus: true, returnStdout: true).trim()
                    if (pidFile) {
                        sh "kill -TERM $pidFile" // Terminate the previous script gracefully
                    }

                    // Start the new script
                    sh 'nohup ./run.sh > /dev/null 2>&1 &'
                    echo $! > /var/run/run.pid
                }
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

