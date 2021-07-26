import QtQuick 2.12
import Ubuntu.Components 1.3
import Ubuntu.Components.Popups 1.3


import "../components"

Page {
    id: page
    
    property var cert
    
    header: PageHeader {
        id: header
        title: cert ? (cert.givenName + " " + cert.familyName) : ""
        trailingActionBar {
            numberOfSlots: 1
            actions: [
                Action {
                    text: "Löschen"
                    iconName: "delete"
                    onTriggered: {
                        var popup = PopupUtils.open(removeCertificate, root);
                        popup.accepted.connect(function() {
                            if (cert) {
                                myapp.removeCert(cert.id)
                                pageStack.removePages(page)
                            }
                        })

                    }
                }
            ]
        }
    }

    ListView {
        width: parent.width
        height: parent.height - header.height - units.gu(2)
        anchors.top: header.bottom
        model: cert ? cert.data.size : 0
        delegate: ListItem {
            height: row.height
            property var item: cert.data.get(index)
            Column {
                leftPadding: units.gu(2)
                rightPadding: units.gu(2)
                topPadding: units.gu(1)
                bottomPadding: units.gu(1)
                spacing: units.gu(1) / 2
                id: row
                width: parent.width
                Label {
                    width: parent.width - parent.leftPadding - parent.rightPadding
                    text: item.title
                    wrapMode: Text.WordWrap
                    textSize: Label.Small
                    opacity: .75
                }
                Label {
                    width: parent.width - parent.leftPadding - parent.rightPadding
                    text: item.subtitle || " "
                    wrapMode: Text.WordWrap
                }
            }
        }
    }

    RemoveDialog {
        id: removeCertificate
    }

}