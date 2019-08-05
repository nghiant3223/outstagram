import React from 'react';
import { Icon, Segment, Grid, Divider, Header, Popup, Button } from 'semantic-ui-react';

import Input from "../../Input/Input";

import "./UploadTypeContainer.css";

function UploadTypeContainer({ expand = true, triggerFileInput, onUrlInputChange, imageURL }) {
    return (
        <Segment placeholder={expand} className={!expand ? "InpandSegment" : ""}>
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
                        <button className="ui button primary" type="button" onClick={triggerFileInput}>{expand ? "Choose your photo" : "Choose another photo"}</button>
                    </Grid.Column>

                    <Grid.Column>
                        {expand && (
                            <Header icon>
                                <Icon name='world' color="black" inverted />
                                Add photo from web
                            </Header>
                        )}

                        <div>
                            <Input width={expand ? "90%" : "80%"} onChange={onUrlInputChange} placeHolder={expand ? "Paste a URL" : "Paste another URL"} value={imageURL} accept="image/*" />
                            <Popup content={`Get image's URL by right-click on it and choose "Copy image address" (for Chrome)`}
                                inverted wide="very" offset="0 5px" position="top center"
                                trigger={<Icon name='question circle outline' size="large" color="blue" className="QuestionMark" />} />
                        </div>
                    </Grid.Column>
                </Grid.Row>
            </Grid>
        </Segment>
    )
}

export default UploadTypeContainer;