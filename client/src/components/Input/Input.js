import React from 'react'

export default function Input({ width, height, value, onChange, placeHolder, accept }) {
    return (
        <div className="ui input" style={{ width, height }}>
            <input placeholder={placeHolder} onChange={onChange} value={value} accept={accept} />
        </div>
    )
}
