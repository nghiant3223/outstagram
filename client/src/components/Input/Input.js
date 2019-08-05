import React from 'react'

export default function Input({ width, height, value, onChange, placeHolder }) {
    return (
        <div className="ui input" style={{ width, height }}>
            <input placeholder={placeHolder} onChange={onChange} value={value} />
        </div>
    )
}
