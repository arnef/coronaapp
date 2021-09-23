import QtQuick 2.12
import Ubuntu.Components 1.3

Rectangle {
    color: cert.color
    width: parent.width
    height: layout.implicitHeight

    signal select(var cert)

    property var cert

    // a little dirty but it works for me
    // property var _textColor: cert.color == "#d2e7fe" ? "black" : "white"

    TapHandler {
        onTapped: {
            select(cert);
        }
    }
    Column {
        id: layout
        height: parent.height
        width: parent.width
        padding: units.gu(2)
        spacing: units.gu(2)
        Row {
            Image {
                width: units.gu(4)
                height: units.gu(4)
                source: "../../assets/" + cert.icon
            }
            Column {
                leftPadding: units.gu(2)
                Label {
                    text: cert.title
                    textSize: Label.Large
                    color: cert.textColor
                }
                Label {
                    text: cert.subTitle
                    color: cert.textColor
                }
            }
        }

        Image {
            property var size: parent.width - units.gu(4)
            source: "image://coronaapp.de.arnef/" + cert.id
            width: size
            height: size
        }

        Item {
            width: layout.width
            height: name.implicitHeight
            anchors.leftMargin: units.gu(2)
            anchors.rightMargin: units.gu(2)
            anchors.left: parent.left
            anchors.right: parent.right
            Label {
                id: name
                text: cert.fullName
                color: cert.textColor
                textSize: Label.Large
            }
            Icon {
                anchors.verticalCenter: parent.verticalCenter
                anchors.right: parent.right
                name: "go-next"
                width: units.gu(2)
                color: cert.textColor
            }
        }
    }
}