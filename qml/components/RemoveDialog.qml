
import QtQuick 2.4
import Ubuntu.Components 1.3
import Ubuntu.Components.Popups 1.3

Component {
    id: removeCertificate
    Dialog {
        id: removeCertificateDialog
        title: "Zertifikat löschen"
        
        signal accepted();

        Button {
            text: "Löschen"
            color: theme.palette.normal.negative
            onClicked: {
                removeCertificateDialog.accepted();
                PopupUtils.close(removeCertificateDialog)
            }
        }

        Button {
            text: "Abbrechen"
            onClicked: {
                PopupUtils.close(removeCertificateDialog)
            }

        }
    }
}
