pipeline {
  agent any
  environment {
    name_final = "sgp-data-history-svc"
  }
  stages {
    stage('Docker Build') {
      agent {
        label 'dev'
      }
      when {
        anyOf {
          branch 'sgp*'
          branch 'sprint-*'
          branch 'master'
        }
      }
      steps {
        script {
          def result = sh(returnStdout: true, script: 'echo "$(docker ps -q --filter name=${name_final})"').trim()
          if (result != "") {
            sh '''
            docker stop ${name_final}
            docker rm -vf ${name_final}
            docker build . -t ${name_final}
            docker system prune -f
	    '''
          } else {
            sh '''
            docker build . -t ${name_final}
            docker system prune -f
	    '''
          }
        }
      }
    }
    stage('SonarQube Analysis') {
      agent {
        label 'dev'
      }
      when {
        anyOf {
          branch 'sgp*'
          branch 'sprint-*'
          branch 'master'
        }
      }
      steps {
        sh 'echo SonarQube'
      }
    }
    stage('Deploy to DEV') {
      agent {
        label 'dev'
      }
      when {
        anyOf {
          branch 'sgp*'
          branch 'sprint-*'
          branch 'master'
        }
      }
      steps {
        script {
          sh '''
          docker run -dt -p 30005:90 --name ${name_final} ${name_final}
          docker system prune -f
	  '''
        }
      }
    }
    stage('Cucumber Tests DEV') {
      agent {
        label 'dev'
      }
      when {
        anyOf {
          branch 'sgp*'
          branch 'sprint-*'
          branch 'master'
        }
      }
      steps {
        echo 'Cucumber Tests'
      }
    }
    stage('Deploy to QA') {
      agent {
        label 'qa'
      }
      when {
        anyOf {
          branch 'sprint-*'
          branch 'master'
        }
      }
      steps {
        script {
          def result = sh(returnStdout: true, script: 'echo "$(docker ps -q --filter name=${name_final})"').trim()
          if (result != "") {
            sh '''
            docker stop ${name_final}
            docker rm -vf ${name_final}
            docker build . -t ${name_final}
            docker run -dt -p 30105:90 --name ${name_final} ${name_final}
            docker system prune -f
	    '''
          } else {
            sh '''
            docker build . -t ${name_final}
            docker run -dt -p 30105:90 --name ${name_final} ${name_final}
            docker system prune -f
	    '''
          }
        }
      }
    }
    stage('QA Approval') {
      agent {
        label 'prd'
      }
      when {
          branch 'master'
      }
      steps {
        input "Aprobacion Tester QA"
      }
    }
    stage('Deploy to PRD') {
      agent {
        label 'prd'
      }
      when {
          branch 'master'
      }
      steps {
        script {
          def result = sh(returnStdout: true, script: 'echo "$(docker ps -q --filter name=${name_final})"').trim()
          if (result != "") {
            sh '''
            docker stop ${name_final}
            docker rm -vf ${name_final}
            docker build . -t ${name_final}
            docker run -dt -p 30205:90 --name ${name_final} ${name_final}
            docker system prune -f
	    '''
          } else {
            sh '''
            docker build . -t ${name_final}
            docker run -dt -p 30205:90 --name ${name_final} ${name_final}
            docker system prune -f
	    '''
          }
        }
      }
    }
  }
}
