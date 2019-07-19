import React from 'react'
import { Button, Header, Placeholder, Popup } from 'semantic-ui-react'

import * as userServices from "../../services/user.service";

const UserInfoPopup = (props) => {
    const { username, trigger } = props;
    const [data, setData] = React.useState(null)

    return (
        <Popup
            on='click'
            onClose={() => {
                setData(null)
            }}
            onOpen={() => {
                setData(null)
                userServices.getUser(username).then(({ data: { data: user } }) => setData(user)).catch((e) => console.log(e));
            }}
            popperDependencies={[!!data]}
            trigger={trigger}
            wide
        >
            {data === null ? (
                <Placeholder style={{ minWidth: '200px' }}>
                    <Placeholder.Header>
                        <Placeholder.Line />
                        <Placeholder.Line />
                    </Placeholder.Header>
                    <Placeholder.Paragraph>
                        <Placeholder.Line length='medium' />
                        <Placeholder.Line length='short' />
                    </Placeholder.Paragraph>
                </Placeholder>
            ) : (
                    <React.Fragment>
                        <Header as='h2' content={data.fullname} subheader={`${data.followerCount} followers`} />
                        <p>{data.description}</p>
                    </React.Fragment>
                )}
        </Popup>
    )
}

export default UserInfoPopup
