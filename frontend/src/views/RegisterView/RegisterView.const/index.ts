import { cn } from '@bem-react/classname';
import { Roles } from '@views/ProfileView/ProfileView.typings';

export const cnRegistration = cn('Registration');

export const registrationCn = cnRegistration();
export const registrationFormCn = cnRegistration('Form');
export const registrationSubmitCn = cnRegistration('Submit');

export const roleSelectOptions = [
    { value: Roles.Administrator, content: Roles.Administrator },
    { value: Roles.Tutor, content: Roles.Tutor },
    { value: Roles.Student, content: Roles.Student },
];
