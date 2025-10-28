import { View, Text } from 'react-native'
import React, { useState } from 'react'
import LoadingScreen from '../../components/LoadingScreen'

const index = () => {

    const [fetched, setFetched] = useState(false)
    const fetch = (filters) => {
        const url = ""
    }

    return (
        <>
            {!fetched ? (
                <>
                    <LoadingScreen />
                    <Text>fetching</Text>
                </>
            ) : (
                <View>
                    <Text>index</Text>
                </View>
            )}
        </>
    )
}

export default index