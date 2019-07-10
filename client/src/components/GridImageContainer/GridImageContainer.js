import React, { Component } from 'react';
// import { Image, Grid, Row, Col } from 'react-bootstrap';
import { Image, Grid } from 'semantic-ui-react';
import PropTypes from 'prop-types';

import "./GridImageContainer.css";

class GridImageContainer extends Component {
    static defaultProps = {
        images: [],
        hideOverlay: false,
        renderOverlay: () => 'Render text',
        overlayBackgroundColor: '#222222',
        onClickEach: null,
        countFrom: 4
    }

    constructor(props) {
        super(props)

        this.state = {
            modal: false,
            countFrom: props.countFrom > 0 && props.countFrom < 5 ? props.countFrom : 5,
            conditionalRender: false
        };

        this.openModal = this.openModal.bind(this);
        this.onClose = this.onClose.bind(this);

        if (props.countFrom <= 0 || props.countFrom > 5) {
            console.warn('countFrom is limited to 5!')
        }
    }

    openModal(index) {
        const { onClickEach, images } = this.props;

        if (onClickEach) {
            return onClickEach({ src: images[index], index })
        }

        this.setState({ modal: true, url: images[index], index })
    }

    onClose() {
        this.setState({ modal: false })
    }

    renderOne() {
        const { images } = this.props;
        const { countFrom } = this.state;
        const overlay = images.length > countFrom && countFrom == 1 ? this.renderCountOverlay(true) : this.renderOverlay();

        return <Grid>
            <Grid.Row className="no-vertical-padding">
                <Grid.Column className={`border height-two background`} onClick={this.openModal.bind(this, 0)} style={{ background: `url(${images[0]})` }}>
                    {overlay}
                </Grid.Column>
            </Grid.Row>
        </Grid>;
    }

    renderTwo() {
        const { images } = this.props;
        const { countFrom } = this.state;
        const overlay = images.length > countFrom && [2, 3].includes(+countFrom) ? this.renderCountOverlay(true) : this.renderOverlay();
        const conditionalRender = [3, 4].includes(images.length) || images.length > +countFrom && [3, 4].includes(+countFrom);

        return <Grid>
            <Grid.Row columns={2} className="no-vertical-padding">
                <Grid.Column className="border height-two background" onClick={this.openModal.bind(this, conditionalRender ? 1 : 0)} style={{ background: `url(${conditionalRender ? images[1] : images[0]})` }}>
                    {this.renderOverlay()}
                </Grid.Column>
                <Grid.Column className="border height-two background" onClick={this.openModal.bind(this, conditionalRender ? 2 : 1)} style={{ background: `url(${conditionalRender ? images[2] : images[1]})` }}>
                    {overlay}
                </Grid.Column>
            </Grid.Row>
        </Grid>;
    }

    renderThree() {
        const { images } = this.props;
        const { countFrom } = this.state;
        const overlay = !countFrom || countFrom > 5 || images.length > countFrom && [4, 5].includes(+countFrom) ? this.renderCountOverlay(true) : this.renderOverlay(conditionalRender ? 3 : 4);
        const conditionalRender = images.length == 4 || images.length > +countFrom && +countFrom == 4;

        return <Grid>
            <Grid.Row columns={3} className="no-vertical-padding">
                <Grid.Column className="border height-three background" onClick={this.openModal.bind(this, conditionalRender ? 1 : 2)} style={{ background: `url(${conditionalRender ? images[1] : images[2]})` }}>
                    {this.renderOverlay(conditionalRender ? 1 : 2)}
                </Grid.Column>
                <Grid.Column className="border height-three background" onClick={this.openModal.bind(this, conditionalRender ? 2 : 3)} style={{ background: `url(${conditionalRender ? images[2] : images[3]})` }}>
                    {this.renderOverlay(conditionalRender ? 2 : 3)}
                </Grid.Column>
                <Grid.Column className="border height-three background" onClick={this.openModal.bind(this, conditionalRender ? 3 : 4)} style={{ background: `url(${conditionalRender ? images[3] : images[4]})` }}>
                    {overlay}
                </Grid.Column>
            </Grid.Row>
        </Grid>;
    }

    // Return overlay of the image, can combine with this.props.renderOverlay to make overlay dynamically
    renderOverlay(id) {
        return null
    }

    renderCountOverlay(more) {
        const { images } = this.props;
        const { countFrom } = this.state;
        const extra = images.length - (countFrom && countFrom > 5 ? 5 : countFrom);

        return [more && <div key="count" className="cover"></div>, more && <div key="count-sub" className="cover-text" style={{ fontSize: '200%' }}><p>+{extra}</p></div>]
    }

    render() {
        const { countFrom } = this.state;
        const { images } = this.props;
        const imagesToShow = [...images];

        if (countFrom && images.length > countFrom) {
            imagesToShow.length = countFrom;
        }

        return (
            <div className="grid-container">
                {[1, 3, 4].includes(imagesToShow.length) && this.renderOne()}
                {imagesToShow.length >= 2 && imagesToShow.length != 4 && this.renderTwo()}
                {imagesToShow.length >= 4 && this.renderThree()}
            </div>
        )
    }

}

GridImageContainer.propTypes = {
    images: PropTypes.array.isRequired,
    hideOverlay: PropTypes.bool,
    renderOverlay: PropTypes.func,
    overlayBackgroundColor: PropTypes.string,
    onClickEach: PropTypes.func,
    countFrom: PropTypes.number,
};

export default GridImageContainer;