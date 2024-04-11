import {DataGrid, GridColDef} from "@mui/x-data-grid";
import {useQuery} from "@tanstack/react-query";
import axios from "axios";
import {Button, Dialog, DialogActions, DialogContent, DialogTitle} from "@mui/material";
import {useState} from "react";

const Recipes = () => {
    const cols: GridColDef[] = [
        { field: "Id", headerName: "ID", width: 70 },
        { field: "Title", headerName: "Title", width: 125 },
        { field: "Description", headerName: "Description", width: 125 },
        { field: "action", headerName: "Action", width: 120, renderCell: (params) => (
            <>
                <Button onClick={() => {
                    setOpenDialog(true)
                    setSelected(params.row)
                }}>Voir</Button>
            </>
            )
        }
    ]

    const [openDialog, setOpenDialog] = useState<boolean>(false);
    const [selected, setSelected] = useState<any>(null);

    const useRecipes = () => {
        return useQuery({
            queryKey: ['recipes'],
            queryFn: async () => {
                const {data} = await axios.get("http://localhost:8080/recipes")
                console.log(data)
                return data ?? []
            }
        })
    }

    const { data, isLoading } = useRecipes();

    if(isLoading) {
        return <div>Loading...</div>
    }

    return (
        <>
            <div className="w-full h-full">
                <DataGrid
                    getRowId={(row) => row.Id}
                    rows={data}
                    columns={cols}
                />
            </div>
            <Dialog open={openDialog}>
                <DialogTitle>{ selected?.Title } - Etapes</DialogTitle>
                <DialogContent>
                    {selected?.Description} <br/>
                    {selected?.CookingTime} min de cuisson<br/>
                    {selected?.ReceiptSteps?.map((step: any) => (<p key={step.Id}>{step.Content}</p>))}
                </DialogContent>
                <DialogActions>
                    <Button onClick={() => {
                        setOpenDialog(false)
                        setSelected(null)
                    }}>Fermer</Button>
                </DialogActions>
            </Dialog>
        </>
    )
}

export default Recipes;