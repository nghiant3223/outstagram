import React from 'react';
import { Icon, Segment, Grid, Divider, Header } from 'semantic-ui-react';

import Input from "../../Input/Input";

function UploadTypeContainer({ expand = true, triggerFileInput, onUrlInputChange, imageURL }) {
    return (
        <Segment placeholder={expand}>
            <Grid columns={2} stackable textAlign='center'>
                <Divider vertical>Or</Divider>
                <Grid.Row verticalAlign='middle'>
                    <Grid.Column>
                        {expand && (
                            <Header icon>
                                <Icon name='grid layout' />
                                Add new photo
                            </Header>
                        )}
                        <button className="ui button primary" type="button" onClick={triggerFileInput}>Choose your photos</button>
                    </Grid.Column>

                    <Grid.Column>
                        {expand && (
                            <Header icon>
                                <Icon name='world' />
                                Add photo from web
                            </Header>
                        )}

                        <div>
                            <Input width="90%" onChange={onUrlInputChange} placeHolder="Paste a URL" value={imageURL} />
                        </div>
                    </Grid.Column>
                </Grid.Row>
            </Grid>
        </Segment>
    )
}

export default UploadTypeContainer;