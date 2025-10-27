import { View, Text, TouchableOpacity, StyleSheet } from 'react-native'
import React from 'react'

const AdvancedFilterItem = ({ label, value, onPress }) => {
  return (
    <TouchableOpacity style={styles.main} onPress={onPress} activeOpacity={0.7}>
      <View style={styles.content}>
        <Text style={styles.label}>{label}</Text>
        {value && <Text style={styles.value}>{value}</Text>}
      </View>
      <Text style={styles.arrow}>›</Text>
    </TouchableOpacity>
  )
}

const styles = StyleSheet.create({
  main: {
    flexDirection: "row",
    justifyContent: "space-between",
    alignItems: "center",
    paddingVertical: 16,
    borderBottomWidth: 1,
    borderBottomColor: "#e0e0e0"
  },
  content: {
    flex: 1
  },
  label: {
    fontSize: 18,
    fontWeight: "500",
    marginBottom: 4
  },
  value: {
    fontSize: 14,
    color: "#666"
  },
  arrow: {
    fontSize: 28,
    color: "#999",
    marginLeft: 8
  }
})

export default AdvancedFilterItem
