
import QtQuick 2.12
import Ubuntu.Components 1.3
import Ubuntu.Components.Popups 1.3

Component {
    id: removeCertificate
    Dialog {
        id: removeCertificateDialog
        title: R.deleteCert
        
        signal accepted();

        Button {
            text: R.delete
            color: theme.palette.normal.negative
            onClicked: {
                removeCertificateDialog.accepted();
                PopupUtils.close(removeCertificateDialog)
            }
        }

        Button {
            text: R.cancel
            onClicked: {
                PopupUtils.close(removeCertificateDialog)
            }

        }
    }
}
