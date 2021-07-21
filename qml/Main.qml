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

import QtQuick 2.12
import Ubuntu.Components 1.3
//import QtQuick.Controls 2.2
import QtQuick.Layouts 1.3
import Qt.labs.settings 1.0
import QtMultimedia 5.4

import "pages"

MainView {
    id: root
    objectName: 'mainView'
    applicationName: 'coronaapp.de.arnef'
    automaticOrientation: true

    width: units.gu(45)
    height: units.gu(75)

    property var _doScan: false

    
    AdaptivePageLayout {
        id: pageStack
        anchors.fill: parent
        primaryPage: certListPage
        layouts: [
            PageColumnsLayout {
                when: true;
                PageColumn {
                    fillWidth: true
                }
            }
        ]

        CertListPage {
            id: certListPage
            onNavigate: function(pageName, pageParams) {
                if (pageName === "cert_details") {
                    certListPage.pageStack.addPageToNextColumn(certListPage, certDetailsPage, { cert: pageParams })
                }   
            }
        }

        CertDetailsPage {
            id: certDetailsPage
            onOpenVaccination: function(cert) {
                certListPage.pageStack.addPageToNextColumn(certDetailsPage, vaccinationDetailsPage, { cert: cert })
            }
        }

        VaccinationDetailsPage {
            id: vaccinationDetailsPage
        }

    }
        
        
    BottomEdge {
        id: bottomEdge
        height: parent.height - header.height
        preloadContent: true
        contentComponent: QRCodeScanner {
            active: _doScan
            width: bottomEdge.width
            height: bottomEdge.height
            onDone: {
                bottomEdge.collapse()
            }
        }
        onCollapseCompleted: {
            console.log("stop scan")
            _doScan = false;
        }
        onCommitCompleted: {
            _doScan = true;
            console.log("start scan")
        }
    }
}
