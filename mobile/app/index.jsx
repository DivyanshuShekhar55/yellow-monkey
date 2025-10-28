import { Text, TouchableOpacity, View } from "react-native";
import { useRouter } from "expo-router";
import GenderFilterSection from "../components/GenderFilterSection"
import AgeFilterSection from "../components/AgeFilterSection"
import LocationFilterSection from "../components/LocationFilterSection"
import AdvancedFilterSection from "../components/AdvancedFilterSection"
import SearchButton from "../components/SearchButton"

export default function Index() {
  const router = useRouter()



  return (
    <View
      style={{
        flex: 1,
        alignItems: "center",
        paddingHorizontal: 8,
        backgroundColor:"#fff"
      }}
    >
      <GenderFilterSection />
      <AgeFilterSection />
      <LocationFilterSection />
      <AdvancedFilterSection />

      <SearchButton/>

    </View>
  );
}
