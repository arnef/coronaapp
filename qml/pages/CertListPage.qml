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
        height: parent.height - header.height
        width: parent.width
        orientation: ListView.Horizontal
        anchors.top: header.bottom
        model: myapp.certs.size
        snapMode: ListView.SnapToItem
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
}