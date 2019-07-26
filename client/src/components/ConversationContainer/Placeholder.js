import React from 'react';
import { Placeholder } from 'semantic-ui-react';

export default function ConversationPlaceholder() {
    return (
        <div className="MessageInfoContainer" style={{ display: "flex" }}>
            <div className="MessageInfoContainer__Info__Detail"
                style={{ height: 66, display: "flex", flexDirection: "column", justifyContent: "center", flexBasis: "66px" }}>
                <div style={{ display: "flex", justifyContent: "center" }}>
                    <Placeholder style={{ height: "1.5em", width: 150 }}>
                        <Placeholder.Header>
                            <Placeholder.Image />
                        </Placeholder.Header>
                    </Placeholder>
                </div>

                <div style={{ display: "flex", justifyContent: "center", marginTop: "0.5em" }}>
                    <Placeholder style={{ height: "1.5em", width: 100 }}>
                        <Placeholder.Header>
                            <Placeholder.Image />
                        </Placeholder.Header>
                    </Placeholder>
                </div>
            </div>

            <div className="ChatboxContainer" style={{ flexGrow: "1" }}>
                <Placeholder style={{ height: "100%", width: "100%", maxWidth: "100%" }}>
                    <Placeholder.Header>
                        <Placeholder.Image />
                    </Placeholder.Header>
                </Placeholder>
            </div>
        </div>
    )
}
