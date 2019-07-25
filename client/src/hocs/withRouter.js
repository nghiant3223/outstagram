import React from 'react';
import { withRouter } from 'react-router';

export const withRouterInnerRef = (WrappedComponent) => {
    class InnerComponentWithRef extends React.Component {
        render() {
            const { forwardRef, ...rest } = this.props;
            return <WrappedComponent {...rest} ref={forwardRef} />;
        }
    }

    const InnerComponentWithRefWithRouter = withRouter(InnerComponentWithRef, { withRef: true });

    return React.forwardRef((props, ref) => {
        return <InnerComponentWithRefWithRouter {...props} forwardRef={ref} />;
    });
}