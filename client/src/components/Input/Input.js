import React from 'react'

export default function Input({ width, value, onChange, placeHolder }) {
    return (
        <div className="ui input" style={{ width }}>
            <input placeholder={placeHolder} onChange={onChange} value={value} />
        </div>
    )
}
