import QtQuick 2.12
import QtMultimedia 5.0
import Ubuntu.Components 1.3

Page {
    id: self
    property bool active: false

    property var result: scanner.text

    onResultChanged: {
        self.checkValue(self.result)
    }

    signal codeParsed(string value)

    header: PageHeader {
        title: R.scan
        trailingActionBar {
            numberOfSlots: 1
            actions: [
                Action {
                    text: "toggle"
                    iconName: "edit-paste"
                    onTriggered: {
                        self.checkValue(Clipboard.data.text)
                    }
                }
            ]
        }
    }

    function checkValue(text) {
        if (text.startsWith("HC1")) {
            self.codeParsed(text);
        }
    }


    Timer {
        id: captureTimer
        interval: 2000
        repeat: true
        running: active
        onTriggered: {
            scanner.scan(crosshair.x, crosshair.y, crosshair.width, crosshair.height)
        }
        onRunningChanged: {
            if (!running) {
                camera.flash.mode = Camera.FlashOff
                camera.stop()
            } else {
                camera.startAndConfigure()
            }
        }
    }

    Camera {
        id: camera
        focus.focusMode: Camera.FocusContinuous
        focus.focusPointMode: Camera.FocusPointAuto

        function startAndConfigure() {
            start();
            focus.focusMode = Camera.FocusContinuous
            focus.focusPointMode = Camera.FocusPointAuto
        }
    }

    VideoOutput {
        source: camera
        visible: active
        focus: visible
        anchors.fill: parent
        fillMode: Image.PreserveAspectCrop
        orientation: -90

        Rectangle {
            id: crosshair
            anchors.centerIn: parent
            width: parent.width * .8
            height: parent.width * .8
            opacity: .5
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
