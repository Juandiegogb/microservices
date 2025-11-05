pipeline {
    agent any

    stages {

        stage('Build') {
            steps {
                echo "Building"
            }
        }

        stage('Go Dependencies') {
            steps {
                echo 'Descargando dependencias de Go...'
                sh 'go mod download'
            }
        }

        stage('Go Test') {
            steps {
                echo 'Ejecutando pruebas unitarias de Go...'

                sh 'go test ./...'
            }
        }

        stage('Go Build') {
            steps {
                echo 'Compilando binario de Go...'
                sh 'go build -o app_binary ./cmd/main.go'
            }
        }

    post {
        always {
            echo 'Pipeline finalizado.'
        }
        failure {
            echo 'El pipeline falló. Revisar logs.'
        }
    }
}