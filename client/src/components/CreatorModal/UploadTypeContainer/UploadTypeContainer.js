import React from 'react';
import { Icon, Segment, Grid, Divider, Header, Popup, Button } from 'semantic-ui-react';

import Input from "../../Input/Input";

import "./UploadTypeContainer.css";

function UploadTypeContainer({ expand = true, triggerFileInput, onUrlInputChange, imageURL }) {
    return (
        <Segment placeholder={expand}>
            <Grid columns={2} stackable textAlign='center'>
                <Divider vertical>Or</Divider>
                <Grid.Row verticalAlign='middle'>
                    <Grid.Column>
                        {expand && (
                            <Header icon>
                                <Icon name='grid layout' color="black" inverted />
                                Add new photo
                            </Header>
                        )}
                        <button className="ui button primary" type="button" onClick={triggerFileInput}>Choose your photos</button>
                    </Grid.Column>

                    <Grid.Column>
                        {expand && (
                            <Header icon>
                                <Icon name='world' color="black" inverted />
                                Add photo from web
                            </Header>
                        )}

                        <div>
                            <Input width="90%" onChange={onUrlInputChange} placeHolder="Paste a URL" value={imageURL} />
                            <Popup content={`Get photo's URL by right-click on image and choose "Copy image address"`} 
                            inverted wide="very" offset="0 5px" position="top center"
                            trigger={<Icon name='question circle outline' size="large" color="blue" className="QuestionMark"/>} />
                        </div>
                    </Grid.Column>
                </Grid.Row>
            </Grid>
        </Segment>
    )
}

export default UploadTypeContainer;