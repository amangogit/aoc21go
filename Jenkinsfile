pipeline{
    // agent any
    agent {
        docker { image 'golang:1.14' }
    }
    stages{
        stage('Test'){
            steps{
                sh 'echo $PATH'
                sh '''
                    go version
                    ls -lah
                '''
                retry(3){
                    sh 'whoami'
                }
                timeout(time: 3, unit: 'MINUTES'){
                    sh 'echo $USER'
                }
            }
        }
    }
}