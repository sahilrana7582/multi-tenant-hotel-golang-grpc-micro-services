import { createSlice, type PayloadAction } from "@reduxjs/toolkit";
import type { Department } from "../types";


interface DepartmentState {
    departments: Department[];
}


const initialState: DepartmentState = {
    departments: [],
};


const departmentSlice = createSlice({
    name: 'departments',
    initialState,
    reducers: {
      setDepartments: (state, action: PayloadAction<Department[]>) => {
        state.departments = action.payload;
      },
      clearDepartments: (state) => {
        state.departments = [];
      }
    },
});


export const {clearDepartments, setDepartments} = departmentSlice.actions;
export default departmentSlice.reducer;
