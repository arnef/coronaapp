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

import "pages"

MainView {
    id: root
    objectName: 'mainView'
    applicationName: 'coronaapp.de.arnef'

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
        }
    }
        
        
    BottomEdge {
        id: bottomEdge
        height: parent.height
        contentComponent: QRCodeScanner {
            active: _doScan
            width: bottomEdge.width
            height: bottomEdge.height
            onDone: {
                bottomEdge.collapse()
            }
        }
        onCollapseCompleted: {
            _doScan = false
        }
        onCommitCompleted: {
            _doScan = true
            scanner.wait()
        }
    }
}
