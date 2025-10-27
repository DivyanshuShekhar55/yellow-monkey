import { View, Text, StyleSheet, ScrollView } from 'react-native'
import React from 'react'
import AdvancedFilterItem from './AdvancedFilterItem'

const advanced_filters = [
  { label: "Education", value: "Select" },
  { label: "Height", value: "Between 5'2\" and 6'3\"" },
]

const AdvancedFilterSection = () => {
  return (
    <View style={styles.main}>
      <View style={styles.header}>
        <Text style={styles.title}>Advanced filters</Text>
        <Text style={styles.reset}>Reset</Text>
      </View>

      <ScrollView style={styles.filterList}>
        {advanced_filters.map((filter, id) => (
          <AdvancedFilterItem
            key={id}
            label={filter.label}
            value={filter.value}
            onPress={() => console.log(`${filter.label} pressed`)}
          />
        ))}
      </ScrollView>
    </View>
  )
}

const styles = StyleSheet.create({
  main: {
    width: "100%",
    marginBottom: 20
  },
  header: {
    flexDirection: "row",
    justifyContent: "space-between",
    alignItems: "center",
    marginBottom: 16
  },
  title: {
    fontWeight: "600",
    fontSize: 24
  },
  reset: {
    fontSize: 16,
    color: "#d4506b"
  },
  filterList: {
    width: "100%"
  }
})

export default AdvancedFilterSection