import { View, Text, StyleSheet } from 'react-native'
import React from 'react'
import FilterBadge from './FilterBadge'

const gender_filters = [
    "Male",
    "Female",
    "Everyone"
]

const GenderFilterSection = () => {
    return (
        <View style={styles.main}>

            <Text style={{fontWeight:"600", fontSize:24, marginBottom:18}} > Who Do You Want To Date</Text>

            <View style={{flexDirection:"row", flexWrap:"wrap", gap:8, marginBottom:12}}>

            {gender_filters.map((i, id) => {
                return (

                    <FilterBadge text={i} key={id} />
                )
            })}
            </View>
        </View>
    )
}

const styles = StyleSheet.create({
    main: {
        width: "100%",
        display: "flex",
        marginBottom:20
    }
})

export default GenderFilterSection