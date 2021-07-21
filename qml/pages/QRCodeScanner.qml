import QtQuick 2.7
import QtMultimedia 5.4
import Ubuntu.Components 1.3

Page {
    id: page
    property bool active: false
    signal done()
    

    header: PageHeader {
        title: "Scan..."
    }

    Timer {
        id: captureTimer
        interval: 2000
        repeat: true
        running: page.active && !scanner.hasResult
        onTriggered: {
            console.log("trigger Scan")
            scanner.scan()
        }
        onRunningChanged: {
            if (!running && scanner.hasResult) {
                page.done()
                camera.stop()
            } else {
                
                camera.startAndConfigure()
            }
            console.log("timer running", running)
        }
    }

    Camera {
        id: camera
        focus {
            focusMode: Camera.FocusInfinity
        }             
        function startAndConfigure() {
            console.log("startAndConfigure")
            start();
            focus.focusMode = Camera.FocusContinuous
            focus.focusPointMode = Camera.FocusPointAuto
        }
       
    }
    Item {
        width: parent.width
        height: parent.height
        

        VideoOutput {
            source: camera
            visible: page.active
            focus : visible // to receive focus and capture key events when visible
            anchors.fill: parent
            fillMode: Image.PreserveAspectCrop
            orientation: -90
        }
    }
}