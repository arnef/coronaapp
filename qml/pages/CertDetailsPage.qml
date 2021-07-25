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
            property var item: cert.data.get(index)
            ListItemLayout {
                title.text: item.title
                subtitle.text: item.subtitle
            }
        }
    }

    RemoveDialog {
        id: removeCertificate
    }

}