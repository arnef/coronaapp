import QtQuick 2.12
import QtMultimedia 5.4
import Ubuntu.Components 1.3

Page {
    id: page
    property bool active: false
    signal done()
    

    header: PageHeader {
        title: "Scan..."
        trailingActionBar {
            numberOfSlots: 1
            actions: [
                Action {
                    text: "toggle"
                    iconName: "edit-paste"
                    onTriggered: {
                        scanner.handleString(Clipboard.data.text)
                    }
                }
            ]
        }
    }

    Timer {
        id: captureTimer
        interval: 2000
        repeat: true
        running: active && !scanner.hasResult
        onTriggered: {
            scanner.scan(crosshair.x, crosshair.y, crosshair.width, crosshair.height)
        }
        onRunningChanged: {
            if (!running) {
                camera.flash.mode = Camera.FlashOff
                camera.stop()
                if (scanner.hasResult) {
                    page.done()
                }
                
            } else {
                camera.start()
            }
            console.log("timer running", running)
        }
    }

    Camera {
        id: camera
        captureMode: Camera.CaptureViewfinder
        focus {
            focusMode: Camera.FocusContinuous
            focusPointMode: Camera.FocusPointCustom
            customFocusPoint: Qt.point(crosshair.x + crosshair.height/2, crosshair.y + crosshair.width/2)
        }
        imageProcessing {
            contrast: 0.66
            saturation: -0.5
        }
    }

    VideoOutput {
        source: camera
        visible: active
        focus: active // to receive focus and capture key events when visible
        anchors.fill: parent
        fillMode: Image.PreserveAspectCrop
        autoOrientation: true

        Rectangle {
            id: crosshair
            anchors.centerIn: parent
            width: parent.width * .6
            height: parent.width * .6
            opacity: .75
            color: "transparent"
            border {
                width: units.gu(1) / 4
                color: UbuntuColors.graphite
            }
            radius: 4
        }

        Button {
            iconName: camera.flash.mode === Camera.FlashVideoLight ? "flash-off" : "flash-on"
            width: units.gu(6)
            height: units.gu(6)
            anchors.bottom: parent.bottom
            anchors.horizontalCenter: parent.horizontalCenter
            anchors.bottomMargin: units.gu(2)
            color: UbuntuColors.graphite
            onClicked: {   
                camera.flash.mode = (camera.flash.mode === Camera.FlashVideoLight ? Camera.FlashOff : Camera.FlashVideoLight)
            }
        }
    }
}