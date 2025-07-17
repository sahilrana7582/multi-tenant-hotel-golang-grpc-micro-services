
export type AccordionItemBase = {
    name: string;
    description: string;
  };
  


export interface AccordionProps<T extends AccordionItemBase> {
    icon: React.ReactNode
    title : string;
    data: T[];
    width ?: string;
}



export const hotelDepartments: AccordionItemBase[] = [
    {
      name: "Housekeeping",
      description: "Responsible for cleaning and maintaining guest rooms and public areas.",
    },
    {
      name: "Front Desk",
      description: "Handles check-ins, check-outs, and guest services at the reception.",
    },
    {
      name: "Food & Beverage",
      description: "Manages dining, room service, and bar operations for guests.",
    },
  ];