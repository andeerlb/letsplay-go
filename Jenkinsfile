pipeline {
    agent any

    environment {
        APP_ENV = 'env'
    }

    parameters {
        choice(name: 'ENV', choices: ['dev', 'staging', 'prod'], description: 'Choose environment')
    }

    stages {
        stage('Build Docker Image') {
            steps {
                script {
                    env.APP_ENV = params.ENV

                    sh """
                    docker build \
                      --build-arg APP_ENV=${env.APP_ENV} \
                      -t letsplay:${env.APP_ENV} \
                      -f docker/Dockerfile .
                    """
                }
            }
        }

        stage('Test Container') {
            steps {
                sh """
                docker run --rm -e APP_ENV=${env.APP_ENV} letsplay:${env.APP_ENV}
                """
            }
        }
    }

    post {
        always {
            echo "Cleaning up Docker images..."
            sh "docker rmi -f letsplay:${env.APP_ENV} || true"
        }
    }
}
