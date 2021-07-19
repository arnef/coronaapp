/*
 * Copyright (C) 2021  arnef
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; version 3.
 *
 * coronaapp is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

import QtQuick 2.7
import Ubuntu.Components 1.3
//import QtQuick.Controls 2.2
import QtQuick.Layouts 1.3
import Qt.labs.settings 1.0
import QtMultimedia 5.4

MainView {
    id: root
    objectName: 'mainView'
    applicationName: 'coronaapp.de.arnef'
    automaticOrientation: true

    width: units.gu(45)
    height: units.gu(75)

    Page {
        anchors.fill: parent

        header: PageHeader {
            id: header
            title: i18n.tr('Corona App')
        }

        Timer {
            id: captureTimer
            interval: 2000
            repeat: true
            running: !testvar.hasResult
            onTriggered: {
                testvar.scan()
            }
            onRunningChanged: {
                console.log("timer running", running)
            }
        }

        Column {
            Item {
                width: units.gu(45)
                height: units.gu(45)
                Camera {
                    id: camera
                    focus.focusMode: Camera.FocusContinuous
                    
                }

                VideoOutput {
                    source: camera
                    visible: true
                    focus : visible // to receive focus and capture key events when visible
                    anchors.fill: parent
                    fillMode: Image.PreserveAspectCrop
                    orientation: -90
                }

            }
            Label {
                id: qrResult
                text: testvar.hasResult ? testvar.result : "Scanning"
            }
        }
    }

}
