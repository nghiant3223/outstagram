import React from 'react';
import PropTypes from 'prop-types'

import "./AmpImage.css";
import { capitalize } from '../../utils/lang';

function AmpImage(props) {
    const { fit, src, container } = props;

    const style = {};
    if (typeof container === "object") {
        style.width = container.width;
        style.height = container.height;
    }

    return (
        <div style={{ ...style }}
            className={container === "auto" ? "AutoResizeContainer" : "FixedContainer"}
            dangerouslySetInnerHTML={{ __html: `<amp-img class="${capitalize(fit)}" layout="fill" src="${src}"/>` }}>
        </div>
    );
}

AmpImage.propTypes = {
    fit: PropTypes.oneOf(["contain", "cover"]).isRequired,
    src: PropTypes.string.isRequired,
    container: PropTypes.oneOfType([PropTypes.oneOf(["auto"]), PropTypes.shape({ width: PropTypes.number.isRequired, height: PropTypes.number.isRequired })])
}

AmpImage.defaultProps = {
    container: "auto"
}

export default AmpImage;