import { View, Text, TouchableOpacity } from 'react-native'
import React from 'react'
import { useRouter } from 'expo-router'


const SearchButton = () => {
    const router = useRouter();

    const onclick = () => {
        console.log("clik")
        router.push('/Matched')
    }
    
    return (
        <TouchableOpacity style={{
            padding: 12,
            borderRadius: 14,
            backgroundColor: '#d4b',
            position: "absolute",
            bottom: 30,
            width: '90%',
            alignItems: "center",
            justifyContent: "center",

        }}
            activeOpacity={0.6}
            onPress={onclick}
        >
            <Text style={{ color: "#fff", fontSize: 18 }}>Search</Text>
        </TouchableOpacity>
    )
}

export default SearchButton