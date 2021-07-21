import QtQuick 2.12
import Ubuntu.Components 1.3


Page {
    id: page
    
    property var cert

    signal openVaccination(var cert)

    header: PageHeader {
        id: header
        title: cert ? (cert.givenName + " " + cert.familyName) : ""
    }

    Column {
        anchors.top: header.bottom
        width: parent.width
        Row {
            leftPadding: units.gu(2)
            topPadding: units.gu(2)
            Label {
                text: "Persönliche Daten"
                textSize: Label.Large
            }
        }

        ListItem {
            height: layout.height
            ListItemLayout {
                id: layout
                title.text: "Name, Vorname / Name, first name"
                subtitle.text: cert ? cert.familyName + ", " + cert.givenName : "-"
            }

        }

        ListItem {
            divider.visible: false
            height: layout2.height
            ListItemLayout {
                id: layout2
                title.text: "Geburtsdatum / Date of birth (YYYY-MM-DD)"
                subtitle.text: cert ? cert.dateOfBirth : "-"
            }
        }

        Row {
            leftPadding: units.gu(2)
            topPadding: units.gu(2)
            Label {
                text: "EU-Zertifikate"
                textSize: Label.Large
            }
        }
        
        Repeater {
            model: cert ? cert.vaccinationCerts.size : 0
            delegate: ListItem {
                property var item: cert.vaccinationCerts.get(index)
                divider.visible: index < (cert.vaccinationCerts.size-1)
                height: layout.height
                onClicked: {
                    console.log(item);
                    openVaccination(item);
                }
                ListItemLayout {
                    id: layout
                    title.text: "Impfzertifikat"
                    subtitle.text: "Impfung " + item.doses + " von " + item.doseSeries
                    summary.text: "Geimpft am " + item.vaccinatedOn
                }
            }
        }
    }
}