import { View, Text, TouchableOpacity, Dimensions, StyleSheet } from 'react-native'
import React from 'react'

const WIDTH = Dimensions.get("screen").width

const FilterBadge = ({text}) => {
  return (
    <TouchableOpacity style={styles.main} activeOpacity={0.6}>
      <Text style={styles.text}>{text}</Text>
    </TouchableOpacity>
  )
}

const styles = StyleSheet.create({
    main :{
        width:WIDTH/4,
        maxWidth:120,
        backgroundColor:"#e0e0e0",
        borderRadius:WIDTH/5,
        paddingVertical:4,
        paddingHorizontal:4
    },
    text:{
        fontSize:18,
        textAlign:"center"
    }
})

export default FilterBadge