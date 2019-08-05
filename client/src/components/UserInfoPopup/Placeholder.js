import React from 'react';
import { Placeholder } from 'semantic-ui-react';

export default function UserInfoPopUpPlaceholder() {
    return (
        <div className="UserInfoPopUp__Container">
            <Placeholder>
                <Placeholder.Image style={{ height: "160px" }} />
            </Placeholder>
            <div style={{ marginLeft: "10%", display: "flex" }}>
                <div>
                    <Placeholder style={{ width: "75px" }}>
                        <Placeholder.Image style={{ height: "75px" }} />
                    </Placeholder>
                </div>

                <div style={{ marginTop: "1em", marginLeft: "1em" }}>
                    <Placeholder style={{ minWidth: '200px' }}>
                        <Placeholder.Header>
                            <Placeholder.Line />
                            <Placeholder.Line length='medium' />
                        </Placeholder.Header>
                    </Placeholder>
                </div>
            </div>

        </div>
    )
}
