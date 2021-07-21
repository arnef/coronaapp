import QtQuick 2.12
import Ubuntu.Components 1.3


Page {
    property var cert

    header: PageHeader {
        id: header
        title: cert ? ("Impfung " + cert.doses + " von " + cert.doseSeries) : ""
    }

    Flickable {
        height: parent.height - header.height
        width: parent.width
        contentHeight: 9 * row0.height
        contentWidth: parent.width
        anchors.top: header.bottom
        Column {
            width: parent.width
            ListItem {
                width: parent.width
                height: row0.height
                ListItemLayout {
                    id: row0
                    title.text: "Zielkrankheit oder -erreger / Disease or agent targeted"
                    subtitle.text: cert ? cert.target : "-"
                }
            }
            ListItem {
                width: parent.width
                height: row1.height
                ListItemLayout {
                    id: row1
                    title.text: "Impfstoff / Vaccine"
                    subtitle.text: cert ? cert.medicalProduct : "-"
                }
            }
            ListItem {
                width: parent.width
                height: row2.height
                ListItemLayout {
                    id: row2
                    title.text: "Art des Impfstoffs / Vaccine Type"
                    subtitle.text: cert ? cert.vaccine : "-"
                }
            }
            ListItem {
                width: parent.width
                height: row3.height
                ListItemLayout {
                    id: row3
                    title.text: "Hersteller / Manufacturer"
                    subtitle.text: cert ? cert.manufacturer : "-"
                }
            }
            ListItem {
                width: parent.width
                height: row4.height
                ListItemLayout {
                    id: row4
                    title.text: "Nummer der Impfung / Number in a series of vaccinations/doses"
                    subtitle.text: cert ? cert.doses + "/" + cert.doseSeries : "-"
                }
            }
            ListItem {
                width: parent.width
                height: row5.height
                ListItemLayout {
                    id: row5
                    title.text: "Datum der Impfung / Date of vaccination (YYYY-MM-DD)"
                    subtitle.text: cert ? cert.vaccinatedOn : "-"
                }
            }
            ListItem {
                width: parent.width
                height: row6.height
                ListItemLayout {
                    id: row6
                    title.text: "Land der Impfung / Member State of vaccination"
                    subtitle.text: cert ? cert.country : "-"
                }
            }
            ListItem {
                width: parent.width
                height: row7.height
                ListItemLayout {
                    id: row7
                    title.text: "ZertifikatausColumnsteller / Certificate issuer"
                    subtitle.text: cert ? cert.issuer : "-"
                }
            }
            ListItem {
                width: parent.width
                height: row8.height
                divider.visible: false
                ListItemLayout {
                    id: row8
                    title.text: "Zertifikatkennung / Unique certificate identifier"
                    subtitle.text: cert ? cert.certificateID : "-"
                }
            }
        }
    }
}