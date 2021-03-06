import QtQuick 2.12
import Ubuntu.Components 1.3

import "../components"

Page {
    id: page
    
    signal navigate(var pageName, var pageParams)

    anchors.fill: parent
    header: PageHeader {
        id: header
        title: i18n.tr("Corona App")
    }

    ListView {
        id: certlist
        // TODO get height from content
        height: parent.width + units.gu(13)
        width: parent.width
        orientation: ListView.Horizontal
        anchors.top: header.bottom
        model: myapp.certs.size
        visible: myapp.certs.size > 0
        snapMode: ListView.SnapOneItem
        highlightRangeMode: ListView.StrictlyEnforceRange  //to update current index, needed for snapMode to work
        delegate: Item {
            width: root.width
            CertItem {
                onSelect: function(cert) {
                    navigate("cert_details", cert)
                }
                cert: myapp.certs.get(index)
            }
        }
    }
    Item {
        visible: myapp.certs.size > 1
        anchors.top: certlist.bottom
        anchors.horizontalCenter: parent.horizontalCenter
        width: dots.implicitWidth
        Row {
            id: dots
            spacing: units.gu(1) / 2
            Repeater {
                model: myapp.certs.size
                delegate: Rectangle {
                    width: units.gu(1)
                    height: units.gu(1)
                    opacity: index === certlist.currentIndex ? 1 : .5
                    color: UbuntuColors.graphite
                    radius: units.gu(1)/2
                }
            }
        }
    }



    Column {
        anchors.centerIn: parent
        anchors.verticalCenterOffset: - (header.height / 2)
        width: parent.width
        padding: units.gu(2)
        spacing: units.gu(2)
        visible: myapp.certs.size == 0
        Image {
            id: icon
            source: "../../assets/no_certs.svg"
            width: units.gu(6)
            height: icon.width
            anchors.horizontalCenter: parent.horizontalCenter
        }
        Label {
            width: parent.width - units.gu(4)
            text: myapp.certs.emptyMessage
            wrapMode: Text.WordWrap
            horizontalAlignment: Text.AlignHCenter
        }
    }
}