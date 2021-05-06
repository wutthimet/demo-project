pipeline {
    environment{
        deployment="demo-project-api"
        namespace="kubernetes-sit"
        registry="ktbirp81.ktb:5000"
        registryCredential='dockerhub'
        //imageName="$namespace/$deployment"
        dockerImage=''
        appversionbuild="0"
        majorbuild="0"
        buildversion = "$appversionbuild.$majorbuild.${BUILD_NUMBER}-SNAPSHOT"
        //integration test
        template = "template-default.hbs"
        FromEmail = "bombeyjubpy@gmail.com"
        ToEmail = "wutthimet@orcsoft.co.th"
    }
    agent any
    stages {
        stage("build") {
            steps {
                echo 'building the application...'
            }
        }
        stage("test") {
            steps {
                echo 'testing the application...'
            }
        }
        stage("deploy") {
            steps {
                echo 'deploying the application...'
            }
        }
        stage('report'){
            agent any
            steps{

                echo 'reporting the application...'

                publishHTML([allowMissing: false, alwaysLinkToLastBuild: false, keepAll: false, reportDir: '', reportFiles: '*.html', reportName: 'HTML Report', reportTitles: ''])
                junit allowEmptyResults: true, testResults:'**/TEST*.xml'
                emailext attachmentsPattern: '**/TEST*.html', from: "${FromEmail}",  body: '''$PROJECT_NAME - Build # $BUILD_NUMBER - $BUILD_STATUS:

                Check console output at $BUILD_URL to view the results.
                Changes:
                $CHANGES

                Changes Since Last Success
                ${CHANGES_SINCE_LAST_SUCCESS}

                Failed Tests:
                ${FAILED_TESTS}

                Total = $TEST_COUNTS
                Failed = ${TEST_COUNTS,var="fail"}

                Total = $TEST_COUNTS
                Passed = ${TEST_COUNTS,var="pass"}

                Build Log:
                ${BUILD_LOG}

                ''', subject: '$PROJECT_NAME - Build # $BUILD_NUMBER - $BUILD_STATUS!', to: "$ToEmail"
            }
        }
    }
}